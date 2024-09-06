package main

// Pcl - Library for 2D/3D image and point cloud processing
// Homepage: https://pointclouds.org/

import (
	"fmt"
	
	"os/exec"
)

func installPcl() {
	// Método 1: Descargar y extraer .tar.gz
	pcl_tar_url := "https://github.com/PointCloudLibrary/pcl/archive/refs/tags/pcl-1.14.1.tar.gz"
	pcl_cmd_tar := exec.Command("curl", "-L", pcl_tar_url, "-o", "package.tar.gz")
	err := pcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcl_zip_url := "https://github.com/PointCloudLibrary/pcl/archive/refs/tags/pcl-1.14.1.zip"
	pcl_cmd_zip := exec.Command("curl", "-L", pcl_zip_url, "-o", "package.zip")
	err = pcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcl_bin_url := "https://github.com/PointCloudLibrary/pcl/archive/refs/tags/pcl-1.14.1.bin"
	pcl_cmd_bin := exec.Command("curl", "-L", pcl_bin_url, "-o", "binary.bin")
	err = pcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcl_src_url := "https://github.com/PointCloudLibrary/pcl/archive/refs/tags/pcl-1.14.1.src.tar.gz"
	pcl_cmd_src := exec.Command("curl", "-L", pcl_src_url, "-o", "source.tar.gz")
	err = pcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcl_cmd_direct := exec.Command("./binary")
	err = pcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cminpack")
exec.Command("latte", "install", "cminpack")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: flann")
exec.Command("latte", "install", "flann")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: libpcap")
exec.Command("latte", "install", "libpcap")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: qhull")
exec.Command("latte", "install", "qhull")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: vtk")
exec.Command("latte", "install", "vtk")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: freeglut")
exec.Command("latte", "install", "freeglut")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
