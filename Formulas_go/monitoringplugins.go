package main

// MonitoringPlugins - Plugins for nagios compatible monitoring systems
// Homepage: https://www.monitoring-plugins.org

import (
	"fmt"
	
	"os/exec"
)

func installMonitoringPlugins() {
	// Método 1: Descargar y extraer .tar.gz
	monitoringplugins_tar_url := "https://www.monitoring-plugins.org/download/monitoring-plugins-2.4.0.tar.gz"
	monitoringplugins_cmd_tar := exec.Command("curl", "-L", monitoringplugins_tar_url, "-o", "package.tar.gz")
	err := monitoringplugins_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	monitoringplugins_zip_url := "https://www.monitoring-plugins.org/download/monitoring-plugins-2.4.0.zip"
	monitoringplugins_cmd_zip := exec.Command("curl", "-L", monitoringplugins_zip_url, "-o", "package.zip")
	err = monitoringplugins_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	monitoringplugins_bin_url := "https://www.monitoring-plugins.org/download/monitoring-plugins-2.4.0.bin"
	monitoringplugins_cmd_bin := exec.Command("curl", "-L", monitoringplugins_bin_url, "-o", "binary.bin")
	err = monitoringplugins_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	monitoringplugins_src_url := "https://www.monitoring-plugins.org/download/monitoring-plugins-2.4.0.src.tar.gz"
	monitoringplugins_cmd_src := exec.Command("curl", "-L", monitoringplugins_src_url, "-o", "source.tar.gz")
	err = monitoringplugins_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	monitoringplugins_cmd_direct := exec.Command("./binary")
	err = monitoringplugins_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: bind")
exec.Command("latte", "install", "bind")
}
