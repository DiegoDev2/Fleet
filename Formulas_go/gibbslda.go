package main

// Gibbslda - Library wrapping imlib2's context API
// Homepage: https://gibbslda.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGibbslda() {
	// Método 1: Descargar y extraer .tar.gz
	gibbslda_tar_url := "https://downloads.sourceforge.net/project/gibbslda/GibbsLDA%2B%2B/0.2/GibbsLDA%2B%2B-0.2.tar.gz"
	gibbslda_cmd_tar := exec.Command("curl", "-L", gibbslda_tar_url, "-o", "package.tar.gz")
	err := gibbslda_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gibbslda_zip_url := "https://downloads.sourceforge.net/project/gibbslda/GibbsLDA%2B%2B/0.2/GibbsLDA%2B%2B-0.2.zip"
	gibbslda_cmd_zip := exec.Command("curl", "-L", gibbslda_zip_url, "-o", "package.zip")
	err = gibbslda_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gibbslda_bin_url := "https://downloads.sourceforge.net/project/gibbslda/GibbsLDA%2B%2B/0.2/GibbsLDA%2B%2B-0.2.bin"
	gibbslda_cmd_bin := exec.Command("curl", "-L", gibbslda_bin_url, "-o", "binary.bin")
	err = gibbslda_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gibbslda_src_url := "https://downloads.sourceforge.net/project/gibbslda/GibbsLDA%2B%2B/0.2/GibbsLDA%2B%2B-0.2.src.tar.gz"
	gibbslda_cmd_src := exec.Command("curl", "-L", gibbslda_src_url, "-o", "source.tar.gz")
	err = gibbslda_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gibbslda_cmd_direct := exec.Command("./binary")
	err = gibbslda_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
