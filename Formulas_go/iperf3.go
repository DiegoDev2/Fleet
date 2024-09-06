package main

// Iperf3 - Update of iperf: measures TCP, UDP, and SCTP bandwidth
// Homepage: https://github.com/esnet/iperf

import (
	"fmt"
	
	"os/exec"
)

func installIperf3() {
	// Método 1: Descargar y extraer .tar.gz
	iperf3_tar_url := "https://downloads.es.net/pub/iperf/iperf-3.17.1.tar.gz"
	iperf3_cmd_tar := exec.Command("curl", "-L", iperf3_tar_url, "-o", "package.tar.gz")
	err := iperf3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iperf3_zip_url := "https://downloads.es.net/pub/iperf/iperf-3.17.1.zip"
	iperf3_cmd_zip := exec.Command("curl", "-L", iperf3_zip_url, "-o", "package.zip")
	err = iperf3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iperf3_bin_url := "https://downloads.es.net/pub/iperf/iperf-3.17.1.bin"
	iperf3_cmd_bin := exec.Command("curl", "-L", iperf3_bin_url, "-o", "binary.bin")
	err = iperf3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iperf3_src_url := "https://downloads.es.net/pub/iperf/iperf-3.17.1.src.tar.gz"
	iperf3_cmd_src := exec.Command("curl", "-L", iperf3_src_url, "-o", "source.tar.gz")
	err = iperf3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iperf3_cmd_direct := exec.Command("./binary")
	err = iperf3_cmd_direct.Run()
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
