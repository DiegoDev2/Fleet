package main

// Dovecot - IMAP/POP3 server
// Homepage: https://dovecot.org/

import (
	"fmt"
	
	"os/exec"
)

func installDovecot() {
	// Método 1: Descargar y extraer .tar.gz
	dovecot_tar_url := "https://dovecot.org/releases/2.3/dovecot-2.3.21.tar.gz"
	dovecot_cmd_tar := exec.Command("curl", "-L", dovecot_tar_url, "-o", "package.tar.gz")
	err := dovecot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dovecot_zip_url := "https://dovecot.org/releases/2.3/dovecot-2.3.21.zip"
	dovecot_cmd_zip := exec.Command("curl", "-L", dovecot_zip_url, "-o", "package.zip")
	err = dovecot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dovecot_bin_url := "https://dovecot.org/releases/2.3/dovecot-2.3.21.bin"
	dovecot_cmd_bin := exec.Command("curl", "-L", dovecot_bin_url, "-o", "binary.bin")
	err = dovecot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dovecot_src_url := "https://dovecot.org/releases/2.3/dovecot-2.3.21.src.tar.gz"
	dovecot_cmd_src := exec.Command("curl", "-L", dovecot_src_url, "-o", "source.tar.gz")
	err = dovecot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dovecot_cmd_direct := exec.Command("./binary")
	err = dovecot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: linux-pam")
	exec.Command("latte", "install", "linux-pam").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
