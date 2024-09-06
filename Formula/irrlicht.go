package main

// Irrlicht - Realtime 3D engine
// Homepage: https://irrlicht.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installIrrlicht() {
	// Método 1: Descargar y extraer .tar.gz
	irrlicht_tar_url := "https://downloads.sourceforge.net/project/irrlicht/Irrlicht%20SDK/1.8/1.8.5/irrlicht-1.8.5.zip"
	irrlicht_cmd_tar := exec.Command("curl", "-L", irrlicht_tar_url, "-o", "package.tar.gz")
	err := irrlicht_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	irrlicht_zip_url := "https://downloads.sourceforge.net/project/irrlicht/Irrlicht%20SDK/1.8/1.8.5/irrlicht-1.8.5.zip"
	irrlicht_cmd_zip := exec.Command("curl", "-L", irrlicht_zip_url, "-o", "package.zip")
	err = irrlicht_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	irrlicht_bin_url := "https://downloads.sourceforge.net/project/irrlicht/Irrlicht%20SDK/1.8/1.8.5/irrlicht-1.8.5.zip"
	irrlicht_cmd_bin := exec.Command("curl", "-L", irrlicht_bin_url, "-o", "binary.bin")
	err = irrlicht_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	irrlicht_src_url := "https://downloads.sourceforge.net/project/irrlicht/Irrlicht%20SDK/1.8/1.8.5/irrlicht-1.8.5.zip"
	irrlicht_cmd_src := exec.Command("curl", "-L", irrlicht_src_url, "-o", "source.tar.gz")
	err = irrlicht_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	irrlicht_cmd_direct := exec.Command("./binary")
	err = irrlicht_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxxf86vm")
	exec.Command("latte", "install", "libxxf86vm").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
