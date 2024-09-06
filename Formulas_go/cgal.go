package main

// Cgal - Computational Geometry Algorithms Library
// Homepage: https://www.cgal.org/

import (
	"fmt"
	
	"os/exec"
)

func installCgal() {
	// Método 1: Descargar y extraer .tar.gz
	cgal_tar_url := "https://github.com/CGAL/cgal/releases/download/v5.6.1/CGAL-5.6.1.tar.xz"
	cgal_cmd_tar := exec.Command("curl", "-L", cgal_tar_url, "-o", "package.tar.gz")
	err := cgal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgal_zip_url := "https://github.com/CGAL/cgal/releases/download/v5.6.1/CGAL-5.6.1.tar.xz"
	cgal_cmd_zip := exec.Command("curl", "-L", cgal_zip_url, "-o", "package.zip")
	err = cgal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgal_bin_url := "https://github.com/CGAL/cgal/releases/download/v5.6.1/CGAL-5.6.1.tar.xz"
	cgal_cmd_bin := exec.Command("curl", "-L", cgal_bin_url, "-o", "binary.bin")
	err = cgal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgal_src_url := "https://github.com/CGAL/cgal/releases/download/v5.6.1/CGAL-5.6.1.tar.xz"
	cgal_cmd_src := exec.Command("curl", "-L", cgal_src_url, "-o", "source.tar.gz")
	err = cgal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgal_cmd_direct := exec.Command("./binary")
	err = cgal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
