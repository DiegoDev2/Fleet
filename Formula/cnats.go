package main

// Cnats - C client for the NATS messaging system
// Homepage: https://github.com/nats-io/nats.c

import (
	"fmt"
	
	"os/exec"
)

func installCnats() {
	// Método 1: Descargar y extraer .tar.gz
	cnats_tar_url := "https://github.com/nats-io/nats.c/archive/refs/tags/v3.8.2.tar.gz"
	cnats_cmd_tar := exec.Command("curl", "-L", cnats_tar_url, "-o", "package.tar.gz")
	err := cnats_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cnats_zip_url := "https://github.com/nats-io/nats.c/archive/refs/tags/v3.8.2.zip"
	cnats_cmd_zip := exec.Command("curl", "-L", cnats_zip_url, "-o", "package.zip")
	err = cnats_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cnats_bin_url := "https://github.com/nats-io/nats.c/archive/refs/tags/v3.8.2.bin"
	cnats_cmd_bin := exec.Command("curl", "-L", cnats_bin_url, "-o", "binary.bin")
	err = cnats_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cnats_src_url := "https://github.com/nats-io/nats.c/archive/refs/tags/v3.8.2.src.tar.gz"
	cnats_cmd_src := exec.Command("curl", "-L", cnats_src_url, "-o", "source.tar.gz")
	err = cnats_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cnats_cmd_direct := exec.Command("./binary")
	err = cnats_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf-c")
	exec.Command("latte", "install", "protobuf-c").Run()
}
