package main

// QuartzWm - XQuartz window-manager
// Homepage: https://gitlab.freedesktop.org/xorg/app/quartz-wm

import (
	"fmt"
	
	"os/exec"
)

func installQuartzWm() {
	// Método 1: Descargar y extraer .tar.gz
	quartzwm_tar_url := "https://gitlab.freedesktop.org/xorg/app/quartz-wm/-/archive/babff9d70f61239c46c53a3e41ce10c7ca1419ce/quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar.bz2"
	quartzwm_cmd_tar := exec.Command("curl", "-L", quartzwm_tar_url, "-o", "package.tar.gz")
	err := quartzwm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quartzwm_zip_url := "https://gitlab.freedesktop.org/xorg/app/quartz-wm/-/archive/babff9d70f61239c46c53a3e41ce10c7ca1419ce/quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar.bz2"
	quartzwm_cmd_zip := exec.Command("curl", "-L", quartzwm_zip_url, "-o", "package.zip")
	err = quartzwm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quartzwm_bin_url := "https://gitlab.freedesktop.org/xorg/app/quartz-wm/-/archive/babff9d70f61239c46c53a3e41ce10c7ca1419ce/quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar.bz2"
	quartzwm_cmd_bin := exec.Command("curl", "-L", quartzwm_bin_url, "-o", "binary.bin")
	err = quartzwm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quartzwm_src_url := "https://gitlab.freedesktop.org/xorg/app/quartz-wm/-/archive/babff9d70f61239c46c53a3e41ce10c7ca1419ce/quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar.bz2"
	quartzwm_cmd_src := exec.Command("curl", "-L", quartzwm_src_url, "-o", "source.tar.gz")
	err = quartzwm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quartzwm_cmd_direct := exec.Command("./binary")
	err = quartzwm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
	fmt.Println("Instalando dependencia: xorg-server")
	exec.Command("latte", "install", "xorg-server").Run()
	fmt.Println("Instalando dependencia: libapplewm")
	exec.Command("latte", "install", "libapplewm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxinerama")
	exec.Command("latte", "install", "libxinerama").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
}
