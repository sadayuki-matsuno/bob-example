{{$table := .Table}}
{{$tAlias := .Aliases.Table $table.Key -}}
{{if .Table.Constraints.Primary -}}
{{$.Importer.Import (printf "github.com/stephenafamo/bob/dialect/%s/im" $.Dialect)}}
{{$.Importer.Import "github.com/stephenafamo/scan" }}



// UpsertByPK uses an executor to upsert the {{$tAlias.UpSingular}}
func (o *{{$tAlias.UpSingular}}) UpsertByPK(ctx context.Context, exec bob.Executor, s *{{$tAlias.UpSingular}}Setter) error {
	columns := s.SetColumns()
  if len(columns) == 0 {
    return nil
  }

  conflictCols := []any{
  {{ range $table.Constraints.Primary.Columns }} 
    "{{.}}",
  {{ end }}
  }

	q := psql.Insert(
		im.Into("{{$table.Key }}"),
	  im.OnConflict(conflictCols...).
		  DoUpdate(im.SetExcluded(columns...)),
		im.Returning(
      {{- range $table.Columns -}}
        "{{- .Name -}}",
      {{- end -}}
    ),
	)

	q.Apply(s)
	ret, err := bob.One(ctx, exec, q, scan.StructMapper[{{$tAlias.UpSingular}}]())
	if err != nil {
		return err
	}
	*o = ret

	return nil
}
{{- end}}

// UpsertDoNothing uses an executor to upsert the {{$tAlias.UpSingular}}
func (o *{{$tAlias.UpSingular}}) UpsertDoNothing(ctx context.Context, exec bob.Executor, s *{{$tAlias.UpSingular}}Setter) error {
  conflictCols := []any{
  {{ range $table.Constraints.Primary.Columns }} 
    "{{.}}",
  {{ end }}
  }

	q := psql.Insert(
		im.Into("{{$table.Key }}"),
		im.Returning(
      {{- range $table.Columns -}}
        "{{- .Name -}}",
      {{- end -}}
    ),
    im.OnConflict(conflictCols...).DoNothing(),
	)

	q.Apply(s)
	ret, err := bob.One(ctx, exec, q, scan.StructMapper[{{$tAlias.UpSingular}}]())
	if err != nil {
		return err
	}
	*o = ret

	return nil
}

{{ range $unique := .Table.Constraints.Uniques -}}
{{$upperUniqueName := titleCase $unique.Name}}
// UpsertBy{{$upperUniqueName}}  uses an executor to upsert the {{$unique.Name}} 
// [{{ range $unique.Columns -}} "{{.}}", {{- end }}]
func (o *{{$tAlias.UpSingular}}) UpsertBy{{$upperUniqueName}}(ctx context.Context, exec bob.Executor, s *{{$tAlias.UpSingular}}Setter) error {
	columns := s.SetColumns()
  if len(columns) == 0 {
    return nil
  }

  conflictCols := []any{
  {{ range $unique.Columns -}} 
    "{{.}}",
  {{- end }}
  }

	q := psql.Insert(
		im.Into("{{$table.Key }}"),
    im.OnConflict(conflictCols...).
  		DoUpdate(im.SetExcluded(columns...)),
		im.Returning(
      {{- range $table.Columns -}}
        "{{- .Name -}}",
      {{- end -}}
    ),
	)

	q.Apply(s)
	ret, err := bob.One(ctx, exec, q, scan.StructMapper[{{$tAlias.UpSingular}}]())
	if err != nil {
		return err
	}
	*o = ret

	return nil
}
{{- end}}
