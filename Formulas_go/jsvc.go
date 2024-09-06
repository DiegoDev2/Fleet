package main

// Jsvc - Wrapper to launch Java applications as daemons
// Homepage: https://commons.apache.org/daemon/jsvc.html

import (
	"fmt"
	
	"os/exec"
)

func installJsvc() {
	// Método 1: Descargar y extraer .tar.gz
	jsvc_tar_url := "https://www.apache.org/dyn/closer.lua?path=commons/daemon/source/commons-daemon-1.4.0-src.tar.gz"
	jsvc_cmd_tar := exec.Command("curl", "-L", jsvc_tar_url, "-o", "package.tar.gz")
	err := jsvc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsvc_zip_url := "https://www.apache.org/dyn/closer.lua?path=commons/daemon/source/commons-daemon-1.4.0-src.zip"
	jsvc_cmd_zip := exec.Command("curl", "-L", jsvc_zip_url, "-o", "package.zip")
	err = jsvc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsvc_bin_url := "https://www.apache.org/dyn/closer.lua?path=commons/daemon/source/commons-daemon-1.4.0-src.bin"
	jsvc_cmd_bin := exec.Command("curl", "-L", jsvc_bin_url, "-o", "binary.bin")
	err = jsvc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsvc_src_url := "https://www.apache.org/dyn/closer.lua?path=commons/daemon/source/commons-daemon-1.4.0-src.src.tar.gz"
	jsvc_cmd_src := exec.Command("curl", "-L", jsvc_src_url, "-o", "source.tar.gz")
	err = jsvc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsvc_cmd_direct := exec.Command("./binary")
	err = jsvc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@21")
exec.Command("latte", "install", "openjdk@21")
}
