{{$table := .Table}}
{{$tAlias := .Aliases.Table $table.Key -}}
{{if .Table.Constraints.Primary -}}
{{$.Importer.Import (printf "github.com/stephenafamo/bob/dialect/%s/im" $.Dialect)}}
{{$.Importer.Import "github.com/stephenafamo/bob/dialect/psql/dialect"}}
{{$.Importer.Import "slices" }}


// UpsertByPK uses an executor to upsert the {{$tAlias.UpSingular}}
func (s {{$tAlias.UpSingular}}Setter) UpsertByPK() bob.Mod[*dialect.InsertQuery] {
  pk := []string{
  {{ range $table.Constraints.Primary.Columns }} 
    "{{.}}",
  {{- end }}
  }

  conflictCols := []any{
  {{ range $table.Constraints.Primary.Columns }} 
    "{{.}}",
  {{- end }}
  }

  return im.OnConflict(conflictCols...).
			DoUpdate(im.SetExcluded(slices.DeleteFunc(s.SetColumns(), func(n string) bool {
					return slices.Contains(pk, n)
		})...))
}

// UpsertDoNothing uses an executor to upsert the {{$tAlias.UpSingular}}
func (s {{$tAlias.UpSingular}}Setter) UpsertDoNothing() bob.Mod[*dialect.InsertQuery] {
  return im.OnConflict().DoNothing()
}

{{ range $unique := .Table.Constraints.Uniques -}}
{{$upperUniqueName := titleCase $unique.Name}}
// UpsertBy{{$upperUniqueName}}  uses an executor to upsert the {{$unique.Name}} 
// [{{ range $unique.Columns -}} "{{.}}", {{- end }}]
func (s {{$tAlias.UpSingular}}Setter) UpsertBy{{$upperUniqueName}}() bob.Mod[*dialect.InsertQuery] {
  pk := []string{
  {{ range $table.Constraints.Primary.Columns }} 
    "{{.}}",
  {{- end }}
  }

  conflictCols := []any{
  {{ range $unique.Columns }} 
    "{{.}}",
  {{- end }}
  }

  return im.OnConflict(conflictCols...).
			DoUpdate(im.SetExcluded(slices.DeleteFunc(s.SetColumns(), func(n string) bool {
					return slices.Contains(pk, n)
		})...))
}
{{- end}}
{{- end}}
