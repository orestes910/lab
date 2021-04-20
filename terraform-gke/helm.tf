resource "helm_release" "kubeview" {

  name       = "kubeview"
  repository = "https://benc-uk.github.io/kubeview/charts"
  chart      = "kubeview"

  values = [
    file("${path.module}/kubeview-values.yaml")
  ]

}

resource "helm_release" "demo-app" {
  
  name       = "demo-app"
  chart      = "${path.module}/demo-app-chart"

  # values = [
  #   file("${path.module}/demo-app-values.yaml")
  # ]

  // Hack to trigger reinstall if chart
  // files are updated by passing in an
  // unused trigger value equal to the
  // null resources output id.
  set {
    name  = "trigger"
    value = tostring(null_resource.chart-update.id)
    type  = "string"
  }
}

resource "null_resource" "chart-update" {
  triggers = {
    chart = sha1(join("", [for f in fileset("${path.module}/demo-app-chart", "**"): filesha1("${path.module}/demo-app-chart/${f}")]))
  }
}