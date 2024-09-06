package main

// Vtk - Toolkit for 3D computer graphics, image processing, and visualization
// Homepage: https://www.vtk.org/

import (
	"fmt"
	
	"os/exec"
)

func installVtk() {
	// Método 1: Descargar y extraer .tar.gz
	vtk_tar_url := "https://www.vtk.org/files/release/9.3/VTK-9.3.1.tar.gz"
	vtk_cmd_tar := exec.Command("curl", "-L", vtk_tar_url, "-o", "package.tar.gz")
	err := vtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vtk_zip_url := "https://www.vtk.org/files/release/9.3/VTK-9.3.1.zip"
	vtk_cmd_zip := exec.Command("curl", "-L", vtk_zip_url, "-o", "package.zip")
	err = vtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vtk_bin_url := "https://www.vtk.org/files/release/9.3/VTK-9.3.1.bin"
	vtk_cmd_bin := exec.Command("curl", "-L", vtk_bin_url, "-o", "binary.bin")
	err = vtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vtk_src_url := "https://www.vtk.org/files/release/9.3/VTK-9.3.1.src.tar.gz"
	vtk_cmd_src := exec.Command("curl", "-L", vtk_src_url, "-o", "source.tar.gz")
	err = vtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vtk_cmd_direct := exec.Command("./binary")
	err = vtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: double-conversion")
	exec.Command("latte", "install", "double-conversion").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: gl2ps")
	exec.Command("latte", "install", "gl2ps").Run()
	fmt.Println("Instalando dependencia: glew")
	exec.Command("latte", "install", "glew").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: jsoncpp")
	exec.Command("latte", "install", "jsoncpp").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: netcdf")
	exec.Command("latte", "install", "netcdf").Run()
	fmt.Println("Instalando dependencia: pugixml")
	exec.Command("latte", "install", "pugixml").Run()
	fmt.Println("Instalando dependencia: pyqt")
	exec.Command("latte", "install", "pyqt").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: theora")
	exec.Command("latte", "install", "theora").Run()
	fmt.Println("Instalando dependencia: utf8cpp")
	exec.Command("latte", "install", "utf8cpp").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
