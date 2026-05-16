{{/*
Expand the name of the chart.
*/}}
{{- define "ks.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "ks.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "ks.chart" -}}
helmChartName: {{ .Chart.Name }}
helmChartVersion: {{.Chart.Version}}
{{- end }}

{{/*
Common labels
*/}}
{{- define "ks.labels" -}}
namespace: {{ .Release.Namespace }}
{{ include "ks.chart" $ }}
releaseName: {{ $.Release.Name }}
managed-by: {{ $.Release.Service }}
{{- end }}

