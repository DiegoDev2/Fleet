package main

// Arabica - XML toolkit written in C++
// Homepage: https://www.jezuk.co.uk/tags/arabica.html

import (
	"fmt"
	
	"os/exec"
)

func installArabica() {
	// Método 1: Descargar y extraer .tar.gz
	arabica_tar_url := "https://github.com/jezhiggins/arabica/archive/refs/tags/2020-April.tar.gz"
	arabica_cmd_tar := exec.Command("curl", "-L", arabica_tar_url, "-o", "package.tar.gz")
	err := arabica_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arabica_zip_url := "https://github.com/jezhiggins/arabica/archive/refs/tags/2020-April.zip"
	arabica_cmd_zip := exec.Command("curl", "-L", arabica_zip_url, "-o", "package.zip")
	err = arabica_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arabica_bin_url := "https://github.com/jezhiggins/arabica/archive/refs/tags/2020-April.bin"
	arabica_cmd_bin := exec.Command("curl", "-L", arabica_bin_url, "-o", "binary.bin")
	err = arabica_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arabica_src_url := "https://github.com/jezhiggins/arabica/archive/refs/tags/2020-April.src.tar.gz"
	arabica_cmd_src := exec.Command("curl", "-L", arabica_src_url, "-o", "source.tar.gz")
	err = arabica_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arabica_cmd_direct := exec.Command("./binary")
	err = arabica_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
