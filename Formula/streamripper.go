package main

// Streamripper - Separate tracks via Shoutcasts title-streaming
// Homepage: https://streamripper.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installStreamripper() {
	// Método 1: Descargar y extraer .tar.gz
	streamripper_tar_url := "https://downloads.sourceforge.net/project/streamripper/streamripper%20%28current%29/1.64.6/streamripper-1.64.6.tar.gz"
	streamripper_cmd_tar := exec.Command("curl", "-L", streamripper_tar_url, "-o", "package.tar.gz")
	err := streamripper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	streamripper_zip_url := "https://downloads.sourceforge.net/project/streamripper/streamripper%20%28current%29/1.64.6/streamripper-1.64.6.zip"
	streamripper_cmd_zip := exec.Command("curl", "-L", streamripper_zip_url, "-o", "package.zip")
	err = streamripper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	streamripper_bin_url := "https://downloads.sourceforge.net/project/streamripper/streamripper%20%28current%29/1.64.6/streamripper-1.64.6.bin"
	streamripper_cmd_bin := exec.Command("curl", "-L", streamripper_bin_url, "-o", "binary.bin")
	err = streamripper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	streamripper_src_url := "https://downloads.sourceforge.net/project/streamripper/streamripper%20%28current%29/1.64.6/streamripper-1.64.6.src.tar.gz"
	streamripper_cmd_src := exec.Command("curl", "-L", streamripper_src_url, "-o", "source.tar.gz")
	err = streamripper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	streamripper_cmd_direct := exec.Command("./binary")
	err = streamripper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
