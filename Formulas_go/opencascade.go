package main

// Opencascade - 3D modeling and numerical simulation software for CAD/CAM/CAE
// Homepage: https://dev.opencascade.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpencascade() {
	// Método 1: Descargar y extraer .tar.gz
	opencascade_tar_url := "https://git.dev.opencascade.org/gitweb/?p=occt.git;a=snapshot;h=refs/tags/V7_8_1;sf=tgz"
	opencascade_cmd_tar := exec.Command("curl", "-L", opencascade_tar_url, "-o", "package.tar.gz")
	err := opencascade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencascade_zip_url := "https://git.dev.opencascade.org/gitweb/?p=occt.git;a=snapshot;h=refs/tags/V7_8_1;sf=tgz"
	opencascade_cmd_zip := exec.Command("curl", "-L", opencascade_zip_url, "-o", "package.zip")
	err = opencascade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencascade_bin_url := "https://git.dev.opencascade.org/gitweb/?p=occt.git;a=snapshot;h=refs/tags/V7_8_1;sf=tgz"
	opencascade_cmd_bin := exec.Command("curl", "-L", opencascade_bin_url, "-o", "binary.bin")
	err = opencascade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencascade_src_url := "https://git.dev.opencascade.org/gitweb/?p=occt.git;a=snapshot;h=refs/tags/V7_8_1;sf=tgz"
	opencascade_cmd_src := exec.Command("curl", "-L", opencascade_src_url, "-o", "source.tar.gz")
	err = opencascade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencascade_cmd_direct := exec.Command("./binary")
	err = opencascade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: rapidjson")
exec.Command("latte", "install", "rapidjson")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freeimage")
exec.Command("latte", "install", "freeimage")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
