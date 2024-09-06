package main

// Openj9 - High performance, scalable, Java virtual machine
// Homepage: https://www.eclipse.org/openj9/

import (
	"fmt"
	
	"os/exec"
)

func installOpenj9() {
	// Método 1: Descargar y extraer .tar.gz
	openj9_tar_url := "https://github.com/eclipse-openj9/openj9.git"
	openj9_cmd_tar := exec.Command("curl", "-L", openj9_tar_url, "-o", "package.tar.gz")
	err := openj9_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openj9_zip_url := "https://github.com/eclipse-openj9/openj9.git"
	openj9_cmd_zip := exec.Command("curl", "-L", openj9_zip_url, "-o", "package.zip")
	err = openj9_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openj9_bin_url := "https://github.com/eclipse-openj9/openj9.git"
	openj9_cmd_bin := exec.Command("curl", "-L", openj9_bin_url, "-o", "binary.bin")
	err = openj9_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openj9_src_url := "https://github.com/eclipse-openj9/openj9.git"
	openj9_cmd_src := exec.Command("curl", "-L", openj9_src_url, "-o", "source.tar.gz")
	err = openj9_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openj9_cmd_direct := exec.Command("./binary")
	err = openj9_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
	fmt.Println("Instalando dependencia: libxtst")
exec.Command("latte", "install", "libxtst")
	fmt.Println("Instalando dependencia: numactl")
exec.Command("latte", "install", "numactl")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
