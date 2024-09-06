package main

// NagiosPlugins - Plugins for the nagios network monitoring system
// Homepage: https://www.nagios-plugins.org/

import (
	"fmt"
	
	"os/exec"
)

func installNagiosPlugins() {
	// Método 1: Descargar y extraer .tar.gz
	nagiosplugins_tar_url := "https://github.com/nagios-plugins/nagios-plugins/releases/download/release-2.4.12/nagios-plugins-2.4.12.tar.gz"
	nagiosplugins_cmd_tar := exec.Command("curl", "-L", nagiosplugins_tar_url, "-o", "package.tar.gz")
	err := nagiosplugins_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nagiosplugins_zip_url := "https://github.com/nagios-plugins/nagios-plugins/releases/download/release-2.4.12/nagios-plugins-2.4.12.zip"
	nagiosplugins_cmd_zip := exec.Command("curl", "-L", nagiosplugins_zip_url, "-o", "package.zip")
	err = nagiosplugins_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nagiosplugins_bin_url := "https://github.com/nagios-plugins/nagios-plugins/releases/download/release-2.4.12/nagios-plugins-2.4.12.bin"
	nagiosplugins_cmd_bin := exec.Command("curl", "-L", nagiosplugins_bin_url, "-o", "binary.bin")
	err = nagiosplugins_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nagiosplugins_src_url := "https://github.com/nagios-plugins/nagios-plugins/releases/download/release-2.4.12/nagios-plugins-2.4.12.src.tar.gz"
	nagiosplugins_cmd_src := exec.Command("curl", "-L", nagiosplugins_src_url, "-o", "source.tar.gz")
	err = nagiosplugins_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nagiosplugins_cmd_direct := exec.Command("./binary")
	err = nagiosplugins_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: bind")
	exec.Command("latte", "install", "bind").Run()
}
