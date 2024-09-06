package main

// Dvdbackup - Rip DVD's from the command-line
// Homepage: https://dvdbackup.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDvdbackup() {
	// Método 1: Descargar y extraer .tar.gz
	dvdbackup_tar_url := "https://downloads.sourceforge.net/project/dvdbackup/dvdbackup/dvdbackup-0.4.2/dvdbackup-0.4.2.tar.gz"
	dvdbackup_cmd_tar := exec.Command("curl", "-L", dvdbackup_tar_url, "-o", "package.tar.gz")
	err := dvdbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvdbackup_zip_url := "https://downloads.sourceforge.net/project/dvdbackup/dvdbackup/dvdbackup-0.4.2/dvdbackup-0.4.2.zip"
	dvdbackup_cmd_zip := exec.Command("curl", "-L", dvdbackup_zip_url, "-o", "package.zip")
	err = dvdbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvdbackup_bin_url := "https://downloads.sourceforge.net/project/dvdbackup/dvdbackup/dvdbackup-0.4.2/dvdbackup-0.4.2.bin"
	dvdbackup_cmd_bin := exec.Command("curl", "-L", dvdbackup_bin_url, "-o", "binary.bin")
	err = dvdbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvdbackup_src_url := "https://downloads.sourceforge.net/project/dvdbackup/dvdbackup/dvdbackup-0.4.2/dvdbackup-0.4.2.src.tar.gz"
	dvdbackup_cmd_src := exec.Command("curl", "-L", dvdbackup_src_url, "-o", "source.tar.gz")
	err = dvdbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvdbackup_cmd_direct := exec.Command("./binary")
	err = dvdbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libdvdread")
	exec.Command("latte", "install", "libdvdread").Run()
}
