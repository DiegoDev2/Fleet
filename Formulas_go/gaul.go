package main

// Gaul - Genetic Algorithm Utility Library
// Homepage: https://gaul.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGaul() {
	// Método 1: Descargar y extraer .tar.gz
	gaul_tar_url := "https://downloads.sourceforge.net/project/gaul/gaul-devel/0.1850-0/gaul-devel-0.1850-0.tar.gz"
	gaul_cmd_tar := exec.Command("curl", "-L", gaul_tar_url, "-o", "package.tar.gz")
	err := gaul_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gaul_zip_url := "https://downloads.sourceforge.net/project/gaul/gaul-devel/0.1850-0/gaul-devel-0.1850-0.zip"
	gaul_cmd_zip := exec.Command("curl", "-L", gaul_zip_url, "-o", "package.zip")
	err = gaul_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gaul_bin_url := "https://downloads.sourceforge.net/project/gaul/gaul-devel/0.1850-0/gaul-devel-0.1850-0.bin"
	gaul_cmd_bin := exec.Command("curl", "-L", gaul_bin_url, "-o", "binary.bin")
	err = gaul_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gaul_src_url := "https://downloads.sourceforge.net/project/gaul/gaul-devel/0.1850-0/gaul-devel-0.1850-0.src.tar.gz"
	gaul_cmd_src := exec.Command("curl", "-L", gaul_src_url, "-o", "source.tar.gz")
	err = gaul_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gaul_cmd_direct := exec.Command("./binary")
	err = gaul_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
