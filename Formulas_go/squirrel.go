package main

// Squirrel - High level, imperative, object-oriented programming language
// Homepage: http://www.squirrel-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installSquirrel() {
	// Método 1: Descargar y extraer .tar.gz
	squirrel_tar_url := "https://downloads.sourceforge.net/project/squirrel/squirrel3/squirrel%203.2%20stable/squirrel_3_2_stable.tar.gz"
	squirrel_cmd_tar := exec.Command("curl", "-L", squirrel_tar_url, "-o", "package.tar.gz")
	err := squirrel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	squirrel_zip_url := "https://downloads.sourceforge.net/project/squirrel/squirrel3/squirrel%203.2%20stable/squirrel_3_2_stable.zip"
	squirrel_cmd_zip := exec.Command("curl", "-L", squirrel_zip_url, "-o", "package.zip")
	err = squirrel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	squirrel_bin_url := "https://downloads.sourceforge.net/project/squirrel/squirrel3/squirrel%203.2%20stable/squirrel_3_2_stable.bin"
	squirrel_cmd_bin := exec.Command("curl", "-L", squirrel_bin_url, "-o", "binary.bin")
	err = squirrel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	squirrel_src_url := "https://downloads.sourceforge.net/project/squirrel/squirrel3/squirrel%203.2%20stable/squirrel_3_2_stable.src.tar.gz"
	squirrel_cmd_src := exec.Command("curl", "-L", squirrel_src_url, "-o", "source.tar.gz")
	err = squirrel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	squirrel_cmd_direct := exec.Command("./binary")
	err = squirrel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
