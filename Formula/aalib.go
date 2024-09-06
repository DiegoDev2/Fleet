package main

// Aalib - Portable ASCII art graphics library
// Homepage: https://aa-project.sourceforge.net/aalib/

import (
	"fmt"
	
	"os/exec"
)

func installAalib() {
	// Método 1: Descargar y extraer .tar.gz
	aalib_tar_url := "https://downloads.sourceforge.net/project/aa-project/aa-lib/1.4rc5/aalib-1.4rc5.tar.gz"
	aalib_cmd_tar := exec.Command("curl", "-L", aalib_tar_url, "-o", "package.tar.gz")
	err := aalib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aalib_zip_url := "https://downloads.sourceforge.net/project/aa-project/aa-lib/1.4rc5/aalib-1.4rc5.zip"
	aalib_cmd_zip := exec.Command("curl", "-L", aalib_zip_url, "-o", "package.zip")
	err = aalib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aalib_bin_url := "https://downloads.sourceforge.net/project/aa-project/aa-lib/1.4rc5/aalib-1.4rc5.bin"
	aalib_cmd_bin := exec.Command("curl", "-L", aalib_bin_url, "-o", "binary.bin")
	err = aalib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aalib_src_url := "https://downloads.sourceforge.net/project/aa-project/aa-lib/1.4rc5/aalib-1.4rc5.src.tar.gz"
	aalib_cmd_src := exec.Command("curl", "-L", aalib_src_url, "-o", "source.tar.gz")
	err = aalib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aalib_cmd_direct := exec.Command("./binary")
	err = aalib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
