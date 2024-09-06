package main

// Sfst - Toolbox for morphological analysers and other FST-based tools
// Homepage: https://www.cis.uni-muenchen.de/~schmid/tools/SFST/

import (
	"fmt"
	
	"os/exec"
)

func installSfst() {
	// Método 1: Descargar y extraer .tar.gz
	sfst_tar_url := "https://www.cis.uni-muenchen.de/~schmid/tools/SFST/data/SFST-1.4.7g.zip"
	sfst_cmd_tar := exec.Command("curl", "-L", sfst_tar_url, "-o", "package.tar.gz")
	err := sfst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sfst_zip_url := "https://www.cis.uni-muenchen.de/~schmid/tools/SFST/data/SFST-1.4.7g.zip"
	sfst_cmd_zip := exec.Command("curl", "-L", sfst_zip_url, "-o", "package.zip")
	err = sfst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sfst_bin_url := "https://www.cis.uni-muenchen.de/~schmid/tools/SFST/data/SFST-1.4.7g.zip"
	sfst_cmd_bin := exec.Command("curl", "-L", sfst_bin_url, "-o", "binary.bin")
	err = sfst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sfst_src_url := "https://www.cis.uni-muenchen.de/~schmid/tools/SFST/data/SFST-1.4.7g.zip"
	sfst_cmd_src := exec.Command("curl", "-L", sfst_src_url, "-o", "source.tar.gz")
	err = sfst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sfst_cmd_direct := exec.Command("./binary")
	err = sfst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
