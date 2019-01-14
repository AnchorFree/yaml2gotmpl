{{ range $key, $value := . }}
   {{ range $login, $role := $value }}
resource "github_membership" "{{ $login }}" {
    username = "{{ $login }}"
    role     = "{{ $role }}"
}
   {{ end }}
{{ end }}
