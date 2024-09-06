package main

// Par - Paragraph reflow for email
// Homepage: http://www.nicemice.net/par/

import (
	"fmt"
	
	"os/exec"
)

func installPar() {
	// Método 1: Descargar y extraer .tar.gz
	par_tar_url := "http://www.nicemice.net/par/Par-1.53.0.tar.gz"
	par_cmd_tar := exec.Command("curl", "-L", par_tar_url, "-o", "package.tar.gz")
	err := par_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	par_zip_url := "http://www.nicemice.net/par/Par-1.53.0.zip"
	par_cmd_zip := exec.Command("curl", "-L", par_zip_url, "-o", "package.zip")
	err = par_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	par_bin_url := "http://www.nicemice.net/par/Par-1.53.0.bin"
	par_cmd_bin := exec.Command("curl", "-L", par_bin_url, "-o", "binary.bin")
	err = par_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	par_src_url := "http://www.nicemice.net/par/Par-1.53.0.src.tar.gz"
	par_cmd_src := exec.Command("curl", "-L", par_src_url, "-o", "source.tar.gz")
	err = par_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	par_cmd_direct := exec.Command("./binary")
	err = par_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
