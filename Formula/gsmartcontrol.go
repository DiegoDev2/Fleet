package main

// Gsmartcontrol - Graphical user interface for smartctl
// Homepage: https://gsmartcontrol.shaduri.dev/

import (
	"fmt"
	
	"os/exec"
)

func installGsmartcontrol() {
	// Método 1: Descargar y extraer .tar.gz
	gsmartcontrol_tar_url := "https://downloads.sourceforge.net/project/gsmartcontrol/1.1.4/gsmartcontrol-1.1.4.tar.bz2"
	gsmartcontrol_cmd_tar := exec.Command("curl", "-L", gsmartcontrol_tar_url, "-o", "package.tar.gz")
	err := gsmartcontrol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gsmartcontrol_zip_url := "https://downloads.sourceforge.net/project/gsmartcontrol/1.1.4/gsmartcontrol-1.1.4.tar.bz2"
	gsmartcontrol_cmd_zip := exec.Command("curl", "-L", gsmartcontrol_zip_url, "-o", "package.zip")
	err = gsmartcontrol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gsmartcontrol_bin_url := "https://downloads.sourceforge.net/project/gsmartcontrol/1.1.4/gsmartcontrol-1.1.4.tar.bz2"
	gsmartcontrol_cmd_bin := exec.Command("curl", "-L", gsmartcontrol_bin_url, "-o", "binary.bin")
	err = gsmartcontrol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gsmartcontrol_src_url := "https://downloads.sourceforge.net/project/gsmartcontrol/1.1.4/gsmartcontrol-1.1.4.tar.bz2"
	gsmartcontrol_cmd_src := exec.Command("curl", "-L", gsmartcontrol_src_url, "-o", "source.tar.gz")
	err = gsmartcontrol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gsmartcontrol_cmd_direct := exec.Command("./binary")
	err = gsmartcontrol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtkmm3")
	exec.Command("latte", "install", "gtkmm3").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: smartmontools")
	exec.Command("latte", "install", "smartmontools").Run()
}
