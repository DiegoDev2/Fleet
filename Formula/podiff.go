package main

// Podiff - Compare textual information in two PO files
// Homepage: https://puszcza.gnu.org.ua/software/podiff/

import (
	"fmt"
	
	"os/exec"
)

func installPodiff() {
	// Método 1: Descargar y extraer .tar.gz
	podiff_tar_url := "https://download.gnu.org.ua/pub/release/podiff/podiff-1.4.tar.gz"
	podiff_cmd_tar := exec.Command("curl", "-L", podiff_tar_url, "-o", "package.tar.gz")
	err := podiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	podiff_zip_url := "https://download.gnu.org.ua/pub/release/podiff/podiff-1.4.zip"
	podiff_cmd_zip := exec.Command("curl", "-L", podiff_zip_url, "-o", "package.zip")
	err = podiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	podiff_bin_url := "https://download.gnu.org.ua/pub/release/podiff/podiff-1.4.bin"
	podiff_cmd_bin := exec.Command("curl", "-L", podiff_bin_url, "-o", "binary.bin")
	err = podiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	podiff_src_url := "https://download.gnu.org.ua/pub/release/podiff/podiff-1.4.src.tar.gz"
	podiff_cmd_src := exec.Command("curl", "-L", podiff_src_url, "-o", "source.tar.gz")
	err = podiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	podiff_cmd_direct := exec.Command("./binary")
	err = podiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
