package main

// Rtabmap - Visual and LiDAR SLAM library and standalone application
// Homepage: https://introlab.github.io/rtabmap

import (
	"fmt"
	
	"os/exec"
)

func installRtabmap() {
	// Método 1: Descargar y extraer .tar.gz
	rtabmap_tar_url := "https://github.com/introlab/rtabmap/archive/refs/tags/0.21.4.tar.gz"
	rtabmap_cmd_tar := exec.Command("curl", "-L", rtabmap_tar_url, "-o", "package.tar.gz")
	err := rtabmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtabmap_zip_url := "https://github.com/introlab/rtabmap/archive/refs/tags/0.21.4.zip"
	rtabmap_cmd_zip := exec.Command("curl", "-L", rtabmap_zip_url, "-o", "package.zip")
	err = rtabmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtabmap_bin_url := "https://github.com/introlab/rtabmap/archive/refs/tags/0.21.4.bin"
	rtabmap_cmd_bin := exec.Command("curl", "-L", rtabmap_bin_url, "-o", "binary.bin")
	err = rtabmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtabmap_src_url := "https://github.com/introlab/rtabmap/archive/refs/tags/0.21.4.src.tar.gz"
	rtabmap_cmd_src := exec.Command("curl", "-L", rtabmap_src_url, "-o", "source.tar.gz")
	err = rtabmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtabmap_cmd_direct := exec.Command("./binary")
	err = rtabmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: g2o")
	exec.Command("latte", "install", "g2o").Run()
	fmt.Println("Instalando dependencia: librealsense")
	exec.Command("latte", "install", "librealsense").Run()
	fmt.Println("Instalando dependencia: octomap")
	exec.Command("latte", "install", "octomap").Run()
	fmt.Println("Instalando dependencia: opencv")
	exec.Command("latte", "install", "opencv").Run()
	fmt.Println("Instalando dependencia: pcl")
	exec.Command("latte", "install", "pcl").Run()
	fmt.Println("Instalando dependencia: pdal")
	exec.Command("latte", "install", "pdal").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: vtk")
	exec.Command("latte", "install", "vtk").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: flann")
	exec.Command("latte", "install", "flann").Run()
	fmt.Println("Instalando dependencia: glew")
	exec.Command("latte", "install", "glew").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
	fmt.Println("Instalando dependencia: libpcap")
	exec.Command("latte", "install", "libpcap").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
}
