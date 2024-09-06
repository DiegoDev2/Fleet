package main

// Libfastjson - Fast json library for C
// Homepage: https://github.com/rsyslog/libfastjson

import (
	"fmt"
	
	"os/exec"
)

func installLibfastjson() {
	// Método 1: Descargar y extraer .tar.gz
	libfastjson_tar_url := "https://download.rsyslog.com/libfastjson/libfastjson-1.2304.0.tar.gz"
	libfastjson_cmd_tar := exec.Command("curl", "-L", libfastjson_tar_url, "-o", "package.tar.gz")
	err := libfastjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfastjson_zip_url := "https://download.rsyslog.com/libfastjson/libfastjson-1.2304.0.zip"
	libfastjson_cmd_zip := exec.Command("curl", "-L", libfastjson_zip_url, "-o", "package.zip")
	err = libfastjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfastjson_bin_url := "https://download.rsyslog.com/libfastjson/libfastjson-1.2304.0.bin"
	libfastjson_cmd_bin := exec.Command("curl", "-L", libfastjson_bin_url, "-o", "binary.bin")
	err = libfastjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfastjson_src_url := "https://download.rsyslog.com/libfastjson/libfastjson-1.2304.0.src.tar.gz"
	libfastjson_cmd_src := exec.Command("curl", "-L", libfastjson_src_url, "-o", "source.tar.gz")
	err = libfastjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfastjson_cmd_direct := exec.Command("./binary")
	err = libfastjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
