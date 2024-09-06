package main

// F3d - Fast and minimalist 3D viewer
// Homepage: https://f3d-app.github.io/f3d/

import (
	"fmt"
	
	"os/exec"
)

func installF3d() {
	// Método 1: Descargar y extraer .tar.gz
	f3d_tar_url := "https://github.com/f3d-app/f3d/archive/refs/tags/v2.5.0.tar.gz"
	f3d_cmd_tar := exec.Command("curl", "-L", f3d_tar_url, "-o", "package.tar.gz")
	err := f3d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	f3d_zip_url := "https://github.com/f3d-app/f3d/archive/refs/tags/v2.5.0.zip"
	f3d_cmd_zip := exec.Command("curl", "-L", f3d_zip_url, "-o", "package.zip")
	err = f3d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	f3d_bin_url := "https://github.com/f3d-app/f3d/archive/refs/tags/v2.5.0.bin"
	f3d_cmd_bin := exec.Command("curl", "-L", f3d_bin_url, "-o", "binary.bin")
	err = f3d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	f3d_src_url := "https://github.com/f3d-app/f3d/archive/refs/tags/v2.5.0.src.tar.gz"
	f3d_cmd_src := exec.Command("curl", "-L", f3d_src_url, "-o", "source.tar.gz")
	err = f3d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	f3d_cmd_direct := exec.Command("./binary")
	err = f3d_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: alembic")
exec.Command("latte", "install", "alembic")
	fmt.Println("Instalando dependencia: assimp")
exec.Command("latte", "install", "assimp")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: opencascade")
exec.Command("latte", "install", "opencascade")
	fmt.Println("Instalando dependencia: vtk")
exec.Command("latte", "install", "vtk")
	fmt.Println("Instalando dependencia: freeimage")
exec.Command("latte", "install", "freeimage")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
