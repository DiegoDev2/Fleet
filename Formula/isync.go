package main

// Isync - Synchronize a maildir with an IMAP server
// Homepage: https://isync.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installIsync() {
	// Método 1: Descargar y extraer .tar.gz
	isync_tar_url := "https://downloads.sourceforge.net/project/isync/isync/1.5.0/isync-1.5.0.tar.gz"
	isync_cmd_tar := exec.Command("curl", "-L", isync_tar_url, "-o", "package.tar.gz")
	err := isync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	isync_zip_url := "https://downloads.sourceforge.net/project/isync/isync/1.5.0/isync-1.5.0.zip"
	isync_cmd_zip := exec.Command("curl", "-L", isync_zip_url, "-o", "package.zip")
	err = isync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	isync_bin_url := "https://downloads.sourceforge.net/project/isync/isync/1.5.0/isync-1.5.0.bin"
	isync_cmd_bin := exec.Command("curl", "-L", isync_bin_url, "-o", "binary.bin")
	err = isync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	isync_src_url := "https://downloads.sourceforge.net/project/isync/isync/1.5.0/isync-1.5.0.src.tar.gz"
	isync_cmd_src := exec.Command("curl", "-L", isync_src_url, "-o", "source.tar.gz")
	err = isync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	isync_cmd_direct := exec.Command("./binary")
	err = isync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
