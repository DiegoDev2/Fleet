package main

// Gobo - Free and portable Eiffel tools and libraries
// Homepage: http://www.gobosoft.com/

import (
	"fmt"
	
	"os/exec"
)

func installGobo() {
	// Método 1: Descargar y extraer .tar.gz
	gobo_tar_url := "https://downloads.sourceforge.net/project/gobo-eiffel/gobo-eiffel/22.01/gobo2201-src.tar.gz"
	gobo_cmd_tar := exec.Command("curl", "-L", gobo_tar_url, "-o", "package.tar.gz")
	err := gobo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gobo_zip_url := "https://downloads.sourceforge.net/project/gobo-eiffel/gobo-eiffel/22.01/gobo2201-src.zip"
	gobo_cmd_zip := exec.Command("curl", "-L", gobo_zip_url, "-o", "package.zip")
	err = gobo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gobo_bin_url := "https://downloads.sourceforge.net/project/gobo-eiffel/gobo-eiffel/22.01/gobo2201-src.bin"
	gobo_cmd_bin := exec.Command("curl", "-L", gobo_bin_url, "-o", "binary.bin")
	err = gobo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gobo_src_url := "https://downloads.sourceforge.net/project/gobo-eiffel/gobo-eiffel/22.01/gobo2201-src.src.tar.gz"
	gobo_cmd_src := exec.Command("curl", "-L", gobo_src_url, "-o", "source.tar.gz")
	err = gobo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gobo_cmd_direct := exec.Command("./binary")
	err = gobo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: eiffelstudio")
exec.Command("latte", "install", "eiffelstudio")
}
