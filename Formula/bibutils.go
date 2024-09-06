package main

// Bibutils - Bibliography conversion utilities
// Homepage: https://sourceforge.net/p/bibutils/home/Bibutils/

import (
	"fmt"
	
	"os/exec"
)

func installBibutils() {
	// Método 1: Descargar y extraer .tar.gz
	bibutils_tar_url := "https://downloads.sourceforge.net/project/bibutils/bibutils_7.2_src.tgz"
	bibutils_cmd_tar := exec.Command("curl", "-L", bibutils_tar_url, "-o", "package.tar.gz")
	err := bibutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bibutils_zip_url := "https://downloads.sourceforge.net/project/bibutils/bibutils_7.2_src.tgz"
	bibutils_cmd_zip := exec.Command("curl", "-L", bibutils_zip_url, "-o", "package.zip")
	err = bibutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bibutils_bin_url := "https://downloads.sourceforge.net/project/bibutils/bibutils_7.2_src.tgz"
	bibutils_cmd_bin := exec.Command("curl", "-L", bibutils_bin_url, "-o", "binary.bin")
	err = bibutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bibutils_src_url := "https://downloads.sourceforge.net/project/bibutils/bibutils_7.2_src.tgz"
	bibutils_cmd_src := exec.Command("curl", "-L", bibutils_src_url, "-o", "source.tar.gz")
	err = bibutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bibutils_cmd_direct := exec.Command("./binary")
	err = bibutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
