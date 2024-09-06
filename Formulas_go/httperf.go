package main

// Httperf - Tool for measuring webserver performance
// Homepage: https://github.com/httperf/httperf

import (
	"fmt"
	
	"os/exec"
)

func installHttperf() {
	// Método 1: Descargar y extraer .tar.gz
	httperf_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/httperf/httperf-0.9.0.tar.gz"
	httperf_cmd_tar := exec.Command("curl", "-L", httperf_tar_url, "-o", "package.tar.gz")
	err := httperf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httperf_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/httperf/httperf-0.9.0.zip"
	httperf_cmd_zip := exec.Command("curl", "-L", httperf_zip_url, "-o", "package.zip")
	err = httperf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httperf_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/httperf/httperf-0.9.0.bin"
	httperf_cmd_bin := exec.Command("curl", "-L", httperf_bin_url, "-o", "binary.bin")
	err = httperf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httperf_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/httperf/httperf-0.9.0.src.tar.gz"
	httperf_cmd_src := exec.Command("curl", "-L", httperf_src_url, "-o", "source.tar.gz")
	err = httperf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httperf_cmd_direct := exec.Command("./binary")
	err = httperf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
