package main

// Openiked - IKEv2 daemon - portable version of OpenBSD iked
// Homepage: https://openiked.org

import (
	"fmt"
	
	"os/exec"
)

func installOpeniked() {
	// Método 1: Descargar y extraer .tar.gz
	openiked_tar_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenIKED/openiked-7.3.tar.gz"
	openiked_cmd_tar := exec.Command("curl", "-L", openiked_tar_url, "-o", "package.tar.gz")
	err := openiked_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openiked_zip_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenIKED/openiked-7.3.zip"
	openiked_cmd_zip := exec.Command("curl", "-L", openiked_zip_url, "-o", "package.zip")
	err = openiked_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openiked_bin_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenIKED/openiked-7.3.bin"
	openiked_cmd_bin := exec.Command("curl", "-L", openiked_bin_url, "-o", "binary.bin")
	err = openiked_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openiked_src_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenIKED/openiked-7.3.src.tar.gz"
	openiked_cmd_src := exec.Command("curl", "-L", openiked_src_url, "-o", "source.tar.gz")
	err = openiked_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openiked_cmd_direct := exec.Command("./binary")
	err = openiked_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
