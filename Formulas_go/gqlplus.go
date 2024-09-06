package main

// Gqlplus - Drop-in replacement for sqlplus, an Oracle SQL client
// Homepage: https://gqlplus.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGqlplus() {
	// Método 1: Descargar y extraer .tar.gz
	gqlplus_tar_url := "https://downloads.sourceforge.net/project/gqlplus/gqlplus/1.16/gqlplus-1.16.tar.gz"
	gqlplus_cmd_tar := exec.Command("curl", "-L", gqlplus_tar_url, "-o", "package.tar.gz")
	err := gqlplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gqlplus_zip_url := "https://downloads.sourceforge.net/project/gqlplus/gqlplus/1.16/gqlplus-1.16.zip"
	gqlplus_cmd_zip := exec.Command("curl", "-L", gqlplus_zip_url, "-o", "package.zip")
	err = gqlplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gqlplus_bin_url := "https://downloads.sourceforge.net/project/gqlplus/gqlplus/1.16/gqlplus-1.16.bin"
	gqlplus_cmd_bin := exec.Command("curl", "-L", gqlplus_bin_url, "-o", "binary.bin")
	err = gqlplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gqlplus_src_url := "https://downloads.sourceforge.net/project/gqlplus/gqlplus/1.16/gqlplus-1.16.src.tar.gz"
	gqlplus_cmd_src := exec.Command("curl", "-L", gqlplus_src_url, "-o", "source.tar.gz")
	err = gqlplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gqlplus_cmd_direct := exec.Command("./binary")
	err = gqlplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
