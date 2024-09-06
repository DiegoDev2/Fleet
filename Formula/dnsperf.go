package main

// Dnsperf - Measure DNS performance by simulating network conditions
// Homepage: https://www.dns-oarc.net/tools/dnsperf

import (
	"fmt"
	
	"os/exec"
)

func installDnsperf() {
	// Método 1: Descargar y extraer .tar.gz
	dnsperf_tar_url := "https://www.dns-oarc.net/files/dnsperf/dnsperf-2.14.0.tar.gz"
	dnsperf_cmd_tar := exec.Command("curl", "-L", dnsperf_tar_url, "-o", "package.tar.gz")
	err := dnsperf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnsperf_zip_url := "https://www.dns-oarc.net/files/dnsperf/dnsperf-2.14.0.zip"
	dnsperf_cmd_zip := exec.Command("curl", "-L", dnsperf_zip_url, "-o", "package.zip")
	err = dnsperf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnsperf_bin_url := "https://www.dns-oarc.net/files/dnsperf/dnsperf-2.14.0.bin"
	dnsperf_cmd_bin := exec.Command("curl", "-L", dnsperf_bin_url, "-o", "binary.bin")
	err = dnsperf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnsperf_src_url := "https://www.dns-oarc.net/files/dnsperf/dnsperf-2.14.0.src.tar.gz"
	dnsperf_cmd_src := exec.Command("curl", "-L", dnsperf_src_url, "-o", "source.tar.gz")
	err = dnsperf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnsperf_cmd_direct := exec.Command("./binary")
	err = dnsperf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: concurrencykit")
	exec.Command("latte", "install", "concurrencykit").Run()
	fmt.Println("Instalando dependencia: ldns")
	exec.Command("latte", "install", "ldns").Run()
	fmt.Println("Instalando dependencia: libnghttp2")
	exec.Command("latte", "install", "libnghttp2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
