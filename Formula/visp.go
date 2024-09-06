package main

// Visp - Visual Servoing Platform library
// Homepage: https://visp.inria.fr/

import (
	"fmt"
	
	"os/exec"
)

func installVisp() {
	// Método 1: Descargar y extraer .tar.gz
	visp_tar_url := "https://visp-doc.inria.fr/download/releases/visp-3.6.0.tar.gz"
	visp_cmd_tar := exec.Command("curl", "-L", visp_tar_url, "-o", "package.tar.gz")
	err := visp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	visp_zip_url := "https://visp-doc.inria.fr/download/releases/visp-3.6.0.zip"
	visp_cmd_zip := exec.Command("curl", "-L", visp_zip_url, "-o", "package.zip")
	err = visp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	visp_bin_url := "https://visp-doc.inria.fr/download/releases/visp-3.6.0.bin"
	visp_cmd_bin := exec.Command("curl", "-L", visp_bin_url, "-o", "binary.bin")
	err = visp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	visp_src_url := "https://visp-doc.inria.fr/download/releases/visp-3.6.0.src.tar.gz"
	visp_cmd_src := exec.Command("curl", "-L", visp_src_url, "-o", "source.tar.gz")
	err = visp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	visp_cmd_direct := exec.Command("./binary")
	err = visp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: gsl")
	exec.Command("latte", "install", "gsl").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libdc1394")
	exec.Command("latte", "install", "libdc1394").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: opencv")
	exec.Command("latte", "install", "opencv").Run()
	fmt.Println("Instalando dependencia: pcl")
	exec.Command("latte", "install", "pcl").Run()
	fmt.Println("Instalando dependencia: vtk")
	exec.Command("latte", "install", "vtk").Run()
	fmt.Println("Instalando dependencia: zbar")
	exec.Command("latte", "install", "zbar").Run()
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
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: libnsl")
	exec.Command("latte", "install", "libnsl").Run()
}
