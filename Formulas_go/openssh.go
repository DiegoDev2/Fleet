package main

// Openssh - OpenBSD freely-licensed SSH connectivity tools
// Homepage: https://www.openssh.com/

import (
	"fmt"
	
	"os/exec"
)

func installOpenssh() {
	// Método 1: Descargar y extraer .tar.gz
	openssh_tar_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.tar.gz"
	openssh_cmd_tar := exec.Command("curl", "-L", openssh_tar_url, "-o", "package.tar.gz")
	err := openssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openssh_zip_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.zip"
	openssh_cmd_zip := exec.Command("curl", "-L", openssh_zip_url, "-o", "package.zip")
	err = openssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openssh_bin_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.bin"
	openssh_cmd_bin := exec.Command("curl", "-L", openssh_bin_url, "-o", "binary.bin")
	err = openssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openssh_src_url := "https://cdn.openbsd.org/pub/OpenBSD/OpenSSH/portable/openssh-9.8p1.src.tar.gz"
	openssh_cmd_src := exec.Command("curl", "-L", openssh_src_url, "-o", "source.tar.gz")
	err = openssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openssh_cmd_direct := exec.Command("./binary")
	err = openssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ldns")
exec.Command("latte", "install", "ldns")
	fmt.Println("Instalando dependencia: libfido2")
exec.Command("latte", "install", "libfido2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
}
