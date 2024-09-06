package main

// Treecc - Aspect-oriented approach to writing compilers
// Homepage: https://gnu.org/software/dotgnu/treecc/treecc.html

import (
	"fmt"
	
	"os/exec"
)

func installTreecc() {
	// Método 1: Descargar y extraer .tar.gz
	treecc_tar_url := "https://download.savannah.gnu.org/releases/dotgnu-pnet/treecc-0.3.10.tar.gz"
	treecc_cmd_tar := exec.Command("curl", "-L", treecc_tar_url, "-o", "package.tar.gz")
	err := treecc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	treecc_zip_url := "https://download.savannah.gnu.org/releases/dotgnu-pnet/treecc-0.3.10.zip"
	treecc_cmd_zip := exec.Command("curl", "-L", treecc_zip_url, "-o", "package.zip")
	err = treecc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	treecc_bin_url := "https://download.savannah.gnu.org/releases/dotgnu-pnet/treecc-0.3.10.bin"
	treecc_cmd_bin := exec.Command("curl", "-L", treecc_bin_url, "-o", "binary.bin")
	err = treecc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	treecc_src_url := "https://download.savannah.gnu.org/releases/dotgnu-pnet/treecc-0.3.10.src.tar.gz"
	treecc_cmd_src := exec.Command("curl", "-L", treecc_src_url, "-o", "source.tar.gz")
	err = treecc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	treecc_cmd_direct := exec.Command("./binary")
	err = treecc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
