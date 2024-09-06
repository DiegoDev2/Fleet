package main

// RiemannClient - C client library for the Riemann monitoring system
// Homepage: https://git.madhouse-project.org/algernon/riemann-c-client

import (
	"fmt"
	
	"os/exec"
)

func installRiemannClient() {
	// Método 1: Descargar y extraer .tar.gz
	riemannclient_tar_url := "https://git.madhouse-project.org/algernon/riemann-c-client/archive/riemann-c-client-2.2.2.tar.gz"
	riemannclient_cmd_tar := exec.Command("curl", "-L", riemannclient_tar_url, "-o", "package.tar.gz")
	err := riemannclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	riemannclient_zip_url := "https://git.madhouse-project.org/algernon/riemann-c-client/archive/riemann-c-client-2.2.2.zip"
	riemannclient_cmd_zip := exec.Command("curl", "-L", riemannclient_zip_url, "-o", "package.zip")
	err = riemannclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	riemannclient_bin_url := "https://git.madhouse-project.org/algernon/riemann-c-client/archive/riemann-c-client-2.2.2.bin"
	riemannclient_cmd_bin := exec.Command("curl", "-L", riemannclient_bin_url, "-o", "binary.bin")
	err = riemannclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	riemannclient_src_url := "https://git.madhouse-project.org/algernon/riemann-c-client/archive/riemann-c-client-2.2.2.src.tar.gz"
	riemannclient_cmd_src := exec.Command("curl", "-L", riemannclient_src_url, "-o", "source.tar.gz")
	err = riemannclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	riemannclient_cmd_direct := exec.Command("./binary")
	err = riemannclient_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf-c")
exec.Command("latte", "install", "protobuf-c")
}
