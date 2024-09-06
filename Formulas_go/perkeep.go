package main

// Perkeep - Lets you permanently keep your stuff, for life
// Homepage: https://perkeep.org/

import (
	"fmt"
	
	"os/exec"
)

func installPerkeep() {
	// Método 1: Descargar y extraer .tar.gz
	perkeep_tar_url := "https://github.com/perkeep/perkeep.git"
	perkeep_cmd_tar := exec.Command("curl", "-L", perkeep_tar_url, "-o", "package.tar.gz")
	err := perkeep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perkeep_zip_url := "https://github.com/perkeep/perkeep.git"
	perkeep_cmd_zip := exec.Command("curl", "-L", perkeep_zip_url, "-o", "package.zip")
	err = perkeep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perkeep_bin_url := "https://github.com/perkeep/perkeep.git"
	perkeep_cmd_bin := exec.Command("curl", "-L", perkeep_bin_url, "-o", "binary.bin")
	err = perkeep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perkeep_src_url := "https://github.com/perkeep/perkeep.git"
	perkeep_cmd_src := exec.Command("curl", "-L", perkeep_src_url, "-o", "source.tar.gz")
	err = perkeep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perkeep_cmd_direct := exec.Command("./binary")
	err = perkeep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.18")
exec.Command("latte", "install", "go@1.18")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
