package html

const EndpointTemplate = `
<div class="endpoint-details-summary">{{.Summary}}</div>
<div class="endpoint-details-method details-http-method-{{.MethodLo}}">{{.MethodUp}}</div>
<div class="endpoint-details-uri">{{.UrlPath}}</div>
<div class="fields-container-title">Parameters</div>
<div class="parameters-description">
  {{if .Parameters}}
    <table>
      <thead>
         <tr>
           <th>Name</th>
           <th>Type</th>
           <th>Mandatory</th>
           <th>Description</th>
         </tr>
      </thead>
      <tbody>
        {{range .Parameters}}
          <tr>
            <td>{{.Name}}</td>
            <td class="json-type-{{.Type}}">{{.Type}}</td>
            <td class="json-mandatory-{{.MandatoryLo}}">{{.Mandatory}}</td>
            <td>{{.Description}}</td>
          </tr>
        {{end}}
      </tbody>
    </table>
  {{else}}
    <div>(none)</div>
  {{end}}
</div>

<div class="fields-container-title">Request body</div>
<div class="parameters-description">
  {{if .RequestBody}}
    <table>
      <thead>
         <tr>
           <th>Name</th>
           <th>Type</th>
           <th>Mandatory</th>
           <th>Description</th>
         </tr>
      </thead>
      <tbody>
        {{range .RequestBody}}
          <tr>
            <td>{{.Name}}</td>
            <td class="json-type-{{.Type}}">{{.Type}}</td>
            <td class="json-mandatory-{{.MandatoryLo}}">{{.Mandatory}}</td>
            <td>{{.Description}}</td>
          </tr>
        {{end}}
      </tbody>
    </table>
  {{else}}
    <div>(none)</div>
  {{end}}
</div>

<div class="fields-container-title">Response body</div>
<div class="parameters-description">
  {{if .ResponseBody}}
    <table>
      <thead>
         <tr>
           <th>Name</th>
           <th>Type</th>
           <th>Mandatory</th>
           <th>Description</th>
         </tr>
      </thead>
      <tbody>
        {{range .ResponseBody}}
          <tr>
            <td>{{.Name}}</td>
            <td class="json-type-{{.Type}}">{{.Type}}</td>
            <td class="json-mandatory-{{.MandatoryLo}}">{{.Mandatory}}</td>
            <td>{{.Description}}</td>
          </tr>
        {{end}}
      </tbody>
    </table>
  {{else}}
    <div>(none)</div>
  {{end}}
</div>

<div class="fields-container-title">Examples</div>
{{range .Examples}}
  <div class="example">
    <div class="example-container">
      <div class="example-summary">{{.Summary}}</div>
      <div class="example-description">{{.Description}}</div>
      <div class="example-request">
        <div class="endpoint-details-method details-http-method-{{.MethodLo}}">{{.Method}}</div>
        <div class="endpoint-details-uri">{{.Uri}}</div>
      </div>
      {{if .RequestBody}}
        <div class="example-request-body"><pre>{{.RequestBody}}</pre></div>
      {{end}}
      <div class="example-response">
        <div class="http-status http-status-{{.StatusCode}}">{{.StatusCode}}</div>
        {{if .ResponseBody}}
          <div class="example-response-body"><pre>{{.ResponseBody}}</pre></div>
        {{end}}
      </div>
    </div>
  </div>
{{end}}
`
