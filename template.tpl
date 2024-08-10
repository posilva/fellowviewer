<!DOCTYPE html>
<html>

<head>
  <title>User Profile</title>
  <style>
    /* Your CSS styles here */
    table {
      border-collapse: collapse;
      width: 100%;
    }

    th,
    td {
      text-align: left;
      padding: 8px;
      border: 1px solid black;
    }
  </style>
</head>

<body>
  <h1>User Information</h1>
  <div class="user-info">
    <p><strong>Email:</strong> {{ .UserInformation.Email }}</p>
    <p><strong>Full Name:</strong> {{ .UserInformation.FullName }}</p>
    <p><strong>First Name:</strong> {{ .UserInformation.FirstName }}</p>
    <p><strong>Last Name:</strong> {{ .UserInformation.LastName }}</p>
    <p><strong>Manager:</strong> {{ .UserInformation.Manager }}</p>
    <p><strong>Title:</strong> {{ .UserInformation.Title }}</p>
    <p><strong>Department:</strong> {{ .UserInformation.Department }}</p>
    <p><strong>Date Joined:</strong> {{ .UserInformation.DateJoined }}</p>
    <p><strong>IP:</strong> {{ .UserInformation.IP }}</p>
    <p><strong>User Agent:</strong> {{ .UserInformation.UserAgent }}</p>
    <p><strong>Last Seen:</strong> {{ .UserInformation.LastSeen }}</p>
  </div>

  <h2>Feedback</h2>
  <ul class="feedback">
    {{ range .Feedback }}
    {{range .}}
    <li>
      <strong>ID:</strong> {{ .ID }}
      <p><strong>Label:</strong> {{ .Label }}</p>
      <p><strong>Value:</strong> {{ .Value }}</p>
    </li>
    {{ end }}
    {{ end }}
  </ul>

  <h2>Calendars</h2>
  <ul class="calendars">
    {{ range .Calendars }}
    <li>
      <h3>{{ .Name }}</h3>
      <ul class="events">
        {{ range .Events }}
        <li>
          <h4>{{ .Title }}</h4>
          <p>
          <pre>{{ unscapeHtml .Description }}</pre>
          </p>
          <p><strong>Location:</strong> {{ .Location }}</p>
          <p><strong>Start:</strong> {{ .Start }}</p>
          <p><strong>End:</strong> {{ .End }}</p>
        </li>
        {{ end }}
      </ul>
    </li>
    {{ end }}
  </ul>

  <h2>Notes</h2>
  <ul class="notes">
    {{ range .Notes }}
    <li>
      <h3>{{ .Title }}</h3>
      <p>{{ handleArray .Content }}</p>
      <p><strong>Start:</strong> {{ .Start }}</p>
      <p><strong>End:</strong> {{ .End }}</p>
    </li>
    {{ end }}
  </ul>

  <h2>Attachments</h2>
  <ul class="attachments">
    {{ range .Attachments }}
    <li><a href="{{ . }}">Download</a></li>
    {{ end }}
  </ul>
</body>

</html>
