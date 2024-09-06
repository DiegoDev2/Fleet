package main

// Appstream - Tools and libraries to work with AppStream metadata
// Homepage: https://www.freedesktop.org/wiki/Distributions/AppStream/

import (
	"fmt"
	
	"os/exec"
)

func installAppstream() {
	// Método 1: Descargar y extraer .tar.gz
	appstream_tar_url := "https://github.com/ximion/appstream/archive/refs/tags/v1.0.3.tar.gz"
	appstream_cmd_tar := exec.Command("curl", "-L", appstream_tar_url, "-o", "package.tar.gz")
	err := appstream_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	appstream_zip_url := "https://github.com/ximion/appstream/archive/refs/tags/v1.0.3.zip"
	appstream_cmd_zip := exec.Command("curl", "-L", appstream_zip_url, "-o", "package.zip")
	err = appstream_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	appstream_bin_url := "https://github.com/ximion/appstream/archive/refs/tags/v1.0.3.bin"
	appstream_cmd_bin := exec.Command("curl", "-L", appstream_bin_url, "-o", "binary.bin")
	err = appstream_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	appstream_src_url := "https://github.com/ximion/appstream/archive/refs/tags/v1.0.3.src.tar.gz"
	appstream_cmd_src := exec.Command("curl", "-L", appstream_src_url, "-o", "source.tar.gz")
	err = appstream_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	appstream_cmd_direct := exec.Command("./binary")
	err = appstream_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: gtk-doc")
	exec.Command("latte", "install", "gtk-doc").Run()
	fmt.Println("Instalando dependencia: itstool")
	exec.Command("latte", "install", "itstool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libxmlb")
	exec.Command("latte", "install", "libxmlb").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gperf")
	exec.Command("latte", "install", "gperf").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
