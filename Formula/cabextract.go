package main

// Cabextract - Extract files from Microsoft cabinet files
// Homepage: https://www.cabextract.org.uk/

import (
	"fmt"
	
	"os/exec"
)

func installCabextract() {
	// Método 1: Descargar y extraer .tar.gz
	cabextract_tar_url := "https://www.cabextract.org.uk/cabextract-1.11.tar.gz"
	cabextract_cmd_tar := exec.Command("curl", "-L", cabextract_tar_url, "-o", "package.tar.gz")
	err := cabextract_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cabextract_zip_url := "https://www.cabextract.org.uk/cabextract-1.11.zip"
	cabextract_cmd_zip := exec.Command("curl", "-L", cabextract_zip_url, "-o", "package.zip")
	err = cabextract_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cabextract_bin_url := "https://www.cabextract.org.uk/cabextract-1.11.bin"
	cabextract_cmd_bin := exec.Command("curl", "-L", cabextract_bin_url, "-o", "binary.bin")
	err = cabextract_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cabextract_src_url := "https://www.cabextract.org.uk/cabextract-1.11.src.tar.gz"
	cabextract_cmd_src := exec.Command("curl", "-L", cabextract_src_url, "-o", "source.tar.gz")
	err = cabextract_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cabextract_cmd_direct := exec.Command("./binary")
	err = cabextract_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
