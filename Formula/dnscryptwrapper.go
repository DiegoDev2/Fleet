package main

// DnscryptWrapper - Server-side proxy that adds dnscrypt support to name resolvers
// Homepage: https://cofyc.github.io/dnscrypt-wrapper/

import (
	"fmt"
	
	"os/exec"
)

func installDnscryptWrapper() {
	// Método 1: Descargar y extraer .tar.gz
	dnscryptwrapper_tar_url := "https://github.com/cofyc/dnscrypt-wrapper/archive/refs/tags/v0.4.2.tar.gz"
	dnscryptwrapper_cmd_tar := exec.Command("curl", "-L", dnscryptwrapper_tar_url, "-o", "package.tar.gz")
	err := dnscryptwrapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dnscryptwrapper_zip_url := "https://github.com/cofyc/dnscrypt-wrapper/archive/refs/tags/v0.4.2.zip"
	dnscryptwrapper_cmd_zip := exec.Command("curl", "-L", dnscryptwrapper_zip_url, "-o", "package.zip")
	err = dnscryptwrapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dnscryptwrapper_bin_url := "https://github.com/cofyc/dnscrypt-wrapper/archive/refs/tags/v0.4.2.bin"
	dnscryptwrapper_cmd_bin := exec.Command("curl", "-L", dnscryptwrapper_bin_url, "-o", "binary.bin")
	err = dnscryptwrapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dnscryptwrapper_src_url := "https://github.com/cofyc/dnscrypt-wrapper/archive/refs/tags/v0.4.2.src.tar.gz"
	dnscryptwrapper_cmd_src := exec.Command("curl", "-L", dnscryptwrapper_src_url, "-o", "source.tar.gz")
	err = dnscryptwrapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dnscryptwrapper_cmd_direct := exec.Command("./binary")
	err = dnscryptwrapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
}
