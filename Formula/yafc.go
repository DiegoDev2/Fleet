package main

// Yafc - Command-line FTP client
// Homepage: https://github.com/sebastinas/yafc

import (
	"fmt"
	
	"os/exec"
)

func installYafc() {
	// Método 1: Descargar y extraer .tar.gz
	yafc_tar_url := "https://deb.debian.org/debian/pool/main/y/yafc/yafc_1.3.7.orig.tar.xz"
	yafc_cmd_tar := exec.Command("curl", "-L", yafc_tar_url, "-o", "package.tar.gz")
	err := yafc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yafc_zip_url := "https://deb.debian.org/debian/pool/main/y/yafc/yafc_1.3.7.orig.tar.xz"
	yafc_cmd_zip := exec.Command("curl", "-L", yafc_zip_url, "-o", "package.zip")
	err = yafc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yafc_bin_url := "https://deb.debian.org/debian/pool/main/y/yafc/yafc_1.3.7.orig.tar.xz"
	yafc_cmd_bin := exec.Command("curl", "-L", yafc_bin_url, "-o", "binary.bin")
	err = yafc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yafc_src_url := "https://deb.debian.org/debian/pool/main/y/yafc/yafc_1.3.7.orig.tar.xz"
	yafc_cmd_src := exec.Command("curl", "-L", yafc_src_url, "-o", "source.tar.gz")
	err = yafc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yafc_cmd_direct := exec.Command("./binary")
	err = yafc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: libbsd")
	exec.Command("latte", "install", "libbsd").Run()
}
