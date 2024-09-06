package main

// Dopewars - No description
// Homepage: https://dopewars.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installDopewars() {
	// Método 1: Descargar y extraer .tar.gz
	dopewars_tar_url := "https://downloads.sourceforge.net/project/dopewars/dopewars/1.6.2/dopewars-1.6.2.tar.gz"
	dopewars_cmd_tar := exec.Command("curl", "-L", dopewars_tar_url, "-o", "package.tar.gz")
	err := dopewars_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dopewars_zip_url := "https://downloads.sourceforge.net/project/dopewars/dopewars/1.6.2/dopewars-1.6.2.zip"
	dopewars_cmd_zip := exec.Command("curl", "-L", dopewars_zip_url, "-o", "package.zip")
	err = dopewars_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dopewars_bin_url := "https://downloads.sourceforge.net/project/dopewars/dopewars/1.6.2/dopewars-1.6.2.bin"
	dopewars_cmd_bin := exec.Command("curl", "-L", dopewars_bin_url, "-o", "binary.bin")
	err = dopewars_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dopewars_src_url := "https://downloads.sourceforge.net/project/dopewars/dopewars/1.6.2/dopewars-1.6.2.src.tar.gz"
	dopewars_cmd_src := exec.Command("curl", "-L", dopewars_src_url, "-o", "source.tar.gz")
	err = dopewars_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dopewars_cmd_direct := exec.Command("./binary")
	err = dopewars_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
