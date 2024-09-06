package main

// Sfcgal - C++ wrapper library around CGAL
// Homepage: https://sfcgal.gitlab.io/SFCGAL/

import (
	"fmt"
	
	"os/exec"
)

func installSfcgal() {
	// Método 1: Descargar y extraer .tar.gz
	sfcgal_tar_url := "https://gitlab.com/sfcgal/SFCGAL/-/archive/v1.5.2/SFCGAL-v1.5.2.tar.gz"
	sfcgal_cmd_tar := exec.Command("curl", "-L", sfcgal_tar_url, "-o", "package.tar.gz")
	err := sfcgal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sfcgal_zip_url := "https://gitlab.com/sfcgal/SFCGAL/-/archive/v1.5.2/SFCGAL-v1.5.2.zip"
	sfcgal_cmd_zip := exec.Command("curl", "-L", sfcgal_zip_url, "-o", "package.zip")
	err = sfcgal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sfcgal_bin_url := "https://gitlab.com/sfcgal/SFCGAL/-/archive/v1.5.2/SFCGAL-v1.5.2.bin"
	sfcgal_cmd_bin := exec.Command("curl", "-L", sfcgal_bin_url, "-o", "binary.bin")
	err = sfcgal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sfcgal_src_url := "https://gitlab.com/sfcgal/SFCGAL/-/archive/v1.5.2/SFCGAL-v1.5.2.src.tar.gz"
	sfcgal_cmd_src := exec.Command("curl", "-L", sfcgal_src_url, "-o", "source.tar.gz")
	err = sfcgal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sfcgal_cmd_direct := exec.Command("./binary")
	err = sfcgal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cgal")
exec.Command("latte", "install", "cgal")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}
