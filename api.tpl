@baseUrl = https://example.com/api


<< range .Service.Groups >>

<<- range .Routes >>

###
         <<- $methodName := .Method ->>
          <<- if eq $methodName "get" >>
            <<- template "method_get" . >>
          <<- else if eq $methodName "post" >>
             <<- template "method_post" . >>
          <<- else if eq $methodName "delete" >>
             <<- template "method_delete" . >>
          <<- end ->>
    << end >>
<< end ->>

<<- define "method_get" >>
# @name << .Handler >> << .JoinedDoc >>
<<toUpper .Method>> {{baseUrl}}<<.Path>> HTTP/1.1
<<- end ->>

<<- define "method_post" >>
# @name << .Handler >> << .JoinedDoc >>
<<toUpper .Method>> {{baseUrl}}<<.Path>> HTTP/1.1
Content-Type: <<contentType .>>

<< genTypes . >>

<<- end ->>

<<- define "method_delete" >>
# @name << .Handler >> << .JoinedDoc >>
<<toUpper .Method>> {{baseUrl}}<<.Path>> HTTP/1.1
<<- end ->>