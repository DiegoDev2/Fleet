package main

// NotmuchMutt - Notmuch integration for Mutt
// Homepage: https://notmuchmail.org/

import (
	"fmt"
	
	"os/exec"
)

func installNotmuchMutt() {
	// Método 1: Descargar y extraer .tar.gz
	notmuchmutt_tar_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuchmutt_cmd_tar := exec.Command("curl", "-L", notmuchmutt_tar_url, "-o", "package.tar.gz")
	err := notmuchmutt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	notmuchmutt_zip_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuchmutt_cmd_zip := exec.Command("curl", "-L", notmuchmutt_zip_url, "-o", "package.zip")
	err = notmuchmutt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	notmuchmutt_bin_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuchmutt_cmd_bin := exec.Command("curl", "-L", notmuchmutt_bin_url, "-o", "binary.bin")
	err = notmuchmutt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	notmuchmutt_src_url := "https://notmuchmail.org/releases/notmuch-0.38.3.tar.xz"
	notmuchmutt_cmd_src := exec.Command("curl", "-L", notmuchmutt_src_url, "-o", "source.tar.gz")
	err = notmuchmutt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	notmuchmutt_cmd_direct := exec.Command("./binary")
	err = notmuchmutt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: notmuch")
	exec.Command("latte", "install", "notmuch").Run()
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
