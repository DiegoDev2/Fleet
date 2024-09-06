package main

// Pymol - Molecular visualization system
// Homepage: https://pymol.org/

import (
	"fmt"
	
	"os/exec"
)

func installPymol() {
	// Método 1: Descargar y extraer .tar.gz
	pymol_tar_url := "https://github.com/schrodinger/pymol-open-source/archive/refs/tags/v3.0.0.tar.gz"
	pymol_cmd_tar := exec.Command("curl", "-L", pymol_tar_url, "-o", "package.tar.gz")
	err := pymol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pymol_zip_url := "https://github.com/schrodinger/pymol-open-source/archive/refs/tags/v3.0.0.zip"
	pymol_cmd_zip := exec.Command("curl", "-L", pymol_zip_url, "-o", "package.zip")
	err = pymol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pymol_bin_url := "https://github.com/schrodinger/pymol-open-source/archive/refs/tags/v3.0.0.bin"
	pymol_cmd_bin := exec.Command("curl", "-L", pymol_bin_url, "-o", "binary.bin")
	err = pymol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pymol_src_url := "https://github.com/schrodinger/pymol-open-source/archive/refs/tags/v3.0.0.src.tar.gz"
	pymol_cmd_src := exec.Command("curl", "-L", pymol_src_url, "-o", "source.tar.gz")
	err = pymol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pymol_cmd_direct := exec.Command("./binary")
	err = pymol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: glm")
exec.Command("latte", "install", "glm")
	fmt.Println("Instalando dependencia: msgpack-cxx")
exec.Command("latte", "install", "msgpack-cxx")
	fmt.Println("Instalando dependencia: sip")
exec.Command("latte", "install", "sip")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pyqt@5")
exec.Command("latte", "install", "pyqt@5")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: freeglut")
exec.Command("latte", "install", "freeglut")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
