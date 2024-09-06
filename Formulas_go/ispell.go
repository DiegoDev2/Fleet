package main

// Ispell - International Ispell
// Homepage: https://www.cs.hmc.edu/~geoff/ispell.html

import (
	"fmt"
	
	"os/exec"
)

func installIspell() {
	// Método 1: Descargar y extraer .tar.gz
	ispell_tar_url := "https://www.cs.hmc.edu/~geoff/tars/ispell-3.4.06.tar.gz"
	ispell_cmd_tar := exec.Command("curl", "-L", ispell_tar_url, "-o", "package.tar.gz")
	err := ispell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ispell_zip_url := "https://www.cs.hmc.edu/~geoff/tars/ispell-3.4.06.zip"
	ispell_cmd_zip := exec.Command("curl", "-L", ispell_zip_url, "-o", "package.zip")
	err = ispell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ispell_bin_url := "https://www.cs.hmc.edu/~geoff/tars/ispell-3.4.06.bin"
	ispell_cmd_bin := exec.Command("curl", "-L", ispell_bin_url, "-o", "binary.bin")
	err = ispell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ispell_src_url := "https://www.cs.hmc.edu/~geoff/tars/ispell-3.4.06.src.tar.gz"
	ispell_cmd_src := exec.Command("curl", "-L", ispell_src_url, "-o", "source.tar.gz")
	err = ispell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ispell_cmd_direct := exec.Command("./binary")
	err = ispell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
