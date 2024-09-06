package main

// Pms - Practical Music Search, an ncurses-based MPD client
// Homepage: https://kimtore.github.io/pms/

import (
	"fmt"
	
	"os/exec"
)

func installPms() {
	// Método 1: Descargar y extraer .tar.gz
	pms_tar_url := "https://downloads.sourceforge.net/project/pms/pms/0.42/pms-0.42.tar.bz2"
	pms_cmd_tar := exec.Command("curl", "-L", pms_tar_url, "-o", "package.tar.gz")
	err := pms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pms_zip_url := "https://downloads.sourceforge.net/project/pms/pms/0.42/pms-0.42.tar.bz2"
	pms_cmd_zip := exec.Command("curl", "-L", pms_zip_url, "-o", "package.zip")
	err = pms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pms_bin_url := "https://downloads.sourceforge.net/project/pms/pms/0.42/pms-0.42.tar.bz2"
	pms_cmd_bin := exec.Command("curl", "-L", pms_bin_url, "-o", "binary.bin")
	err = pms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pms_src_url := "https://downloads.sourceforge.net/project/pms/pms/0.42/pms-0.42.tar.bz2"
	pms_cmd_src := exec.Command("curl", "-L", pms_src_url, "-o", "source.tar.gz")
	err = pms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pms_cmd_direct := exec.Command("./binary")
	err = pms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
