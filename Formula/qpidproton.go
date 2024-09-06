package main

// QpidProton - High-performance, lightweight AMQP 1.0 messaging library
// Homepage: https://qpid.apache.org/proton/

import (
	"fmt"
	
	"os/exec"
)

func installQpidProton() {
	// Método 1: Descargar y extraer .tar.gz
	qpidproton_tar_url := "https://www.apache.org/dyn/closer.lua?path=qpid/proton/0.39.0/qpid-proton-0.39.0.tar.gz"
	qpidproton_cmd_tar := exec.Command("curl", "-L", qpidproton_tar_url, "-o", "package.tar.gz")
	err := qpidproton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qpidproton_zip_url := "https://www.apache.org/dyn/closer.lua?path=qpid/proton/0.39.0/qpid-proton-0.39.0.zip"
	qpidproton_cmd_zip := exec.Command("curl", "-L", qpidproton_zip_url, "-o", "package.zip")
	err = qpidproton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qpidproton_bin_url := "https://www.apache.org/dyn/closer.lua?path=qpid/proton/0.39.0/qpid-proton-0.39.0.bin"
	qpidproton_cmd_bin := exec.Command("curl", "-L", qpidproton_bin_url, "-o", "binary.bin")
	err = qpidproton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qpidproton_src_url := "https://www.apache.org/dyn/closer.lua?path=qpid/proton/0.39.0/qpid-proton-0.39.0.src.tar.gz"
	qpidproton_cmd_src := exec.Command("curl", "-L", qpidproton_src_url, "-o", "source.tar.gz")
	err = qpidproton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qpidproton_cmd_direct := exec.Command("./binary")
	err = qpidproton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
