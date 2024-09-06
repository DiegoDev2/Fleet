package main

// Libssh - C library SSHv1/SSHv2 client and server protocols
// Homepage: https://www.libssh.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibssh() {
	// Método 1: Descargar y extraer .tar.gz
	libssh_tar_url := "https://www.libssh.org/files/0.11/libssh-0.11.1.tar.xz"
	libssh_cmd_tar := exec.Command("curl", "-L", libssh_tar_url, "-o", "package.tar.gz")
	err := libssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libssh_zip_url := "https://www.libssh.org/files/0.11/libssh-0.11.1.tar.xz"
	libssh_cmd_zip := exec.Command("curl", "-L", libssh_zip_url, "-o", "package.zip")
	err = libssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libssh_bin_url := "https://www.libssh.org/files/0.11/libssh-0.11.1.tar.xz"
	libssh_cmd_bin := exec.Command("curl", "-L", libssh_bin_url, "-o", "binary.bin")
	err = libssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libssh_src_url := "https://www.libssh.org/files/0.11/libssh-0.11.1.tar.xz"
	libssh_cmd_src := exec.Command("curl", "-L", libssh_src_url, "-o", "source.tar.gz")
	err = libssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libssh_cmd_direct := exec.Command("./binary")
	err = libssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
