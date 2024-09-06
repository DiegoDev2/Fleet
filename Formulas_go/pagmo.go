package main

// Pagmo - Scientific library for massively parallel optimization
// Homepage: https://esa.github.io/pagmo2/

import (
	"fmt"
	
	"os/exec"
)

func installPagmo() {
	// Método 1: Descargar y extraer .tar.gz
	pagmo_tar_url := "https://github.com/esa/pagmo2/archive/refs/tags/v2.19.1.tar.gz"
	pagmo_cmd_tar := exec.Command("curl", "-L", pagmo_tar_url, "-o", "package.tar.gz")
	err := pagmo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pagmo_zip_url := "https://github.com/esa/pagmo2/archive/refs/tags/v2.19.1.zip"
	pagmo_cmd_zip := exec.Command("curl", "-L", pagmo_zip_url, "-o", "package.zip")
	err = pagmo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pagmo_bin_url := "https://github.com/esa/pagmo2/archive/refs/tags/v2.19.1.bin"
	pagmo_cmd_bin := exec.Command("curl", "-L", pagmo_bin_url, "-o", "binary.bin")
	err = pagmo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pagmo_src_url := "https://github.com/esa/pagmo2/archive/refs/tags/v2.19.1.src.tar.gz"
	pagmo_cmd_src := exec.Command("curl", "-L", pagmo_src_url, "-o", "source.tar.gz")
	err = pagmo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pagmo_cmd_direct := exec.Command("./binary")
	err = pagmo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: nlopt")
exec.Command("latte", "install", "nlopt")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
}
