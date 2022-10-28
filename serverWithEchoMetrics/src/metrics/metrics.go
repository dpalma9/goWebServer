package metrics

//import "fmt"
import (
	//"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus"
)

func Test(myCounter *prometheus.GaugeVec) string {
  // Prometheus
  //myCounter := prometheus.NewGauge(prometheus.GaugeOpts{
  //	Name:        "my_handler_executions",
  //	Help:        "Counts executions of my handler function.",
  //	ConstLabels: prometheus.Labels{"version": "1234"},
  //})
  //if err := prometheus.Register(myCounter); err != nil {
  //	log.Fatal(err)
  //}
  //
  //fmt.Println("Mensaje de prueba")
  var texto string
  texto = "MENSAJE DE PRUEBA"
  //myCounter.Labels("testserver").Set(9.9)
  myCounter.WithLabelValues("testserver").Set(9.9)
  //myCounter.With(prometheus.Labels{"role": "testserver"}).Set(9.9)
  return texto
}
